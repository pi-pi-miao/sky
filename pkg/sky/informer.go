package sky

import (
	"errors"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"
)

func (s *SkyConfig) CreateInformer()error {
	factory := informers.NewSharedInformerFactory(s.Client, 0)
	s.Informer = factory.Core().V1().ConfigMaps()
	s.NamespaceInformer = factory.Core().V1().Namespaces()
	s.Informers = append(s.Informers, s.Informer.Informer().HasSynced, s.NamespaceInformer.Informer().HasSynced)
	go factory.Start(s.Stop)
	if !cache.WaitForCacheSync(s.Stop, s.Informers...) {
		return errors.New("[CreateInformer] timed out waiting for caches to sync")
	}
	s.Informer.Informer().AddEventHandler(s)
	return nil
}

func (s *SkyConfig) CheckNs() error {
	if _, err := s.NamespaceInformer.Lister().Get(s.NameSpace); err != nil {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			_, err := s.Client.CoreV1().Namespaces().Create(&v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       Namespace,
					APIVersion: Version,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      s.NameSpace,
					Namespace: s.NameSpace,
				},
			})
			return err
		})
		return err
	}
	return nil
}

func (s *SkyConfig)CheckUser()(err error){
	if _,err = s.Informer.Lister().ConfigMaps(s.NameSpace).Get(s.User.UserStoreName);err == nil {
		return
	}
	if err = s.createUserDb();err != nil {
		return
	}
	return
}

func (s *SkyConfig) createUserDb() error {
	labels := make(map[string]string, 1)
	for k, _ := range s.Config.UserLabels {
		labels[k] = s.Config.UserLabels[k].Value
	}
	if err := s.Informer.Lister().ConfigMaps(s.NameSpace); err != nil {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			_, err := s.Client.CoreV1().ConfigMaps(s.NameSpace).Create(&v1.ConfigMap{
				TypeMeta: metav1.TypeMeta{
					Kind:       s.Config.User.UserKind,
					APIVersion: s.Config.User.UserVersion,
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      s.User.UserStoreName,
					Namespace: s.NameSpace,
					UID:       uuid.NewUUID(),
					Labels:    labels,
				},
			})
			return err
		})
		return err
	}
	return nil
}
