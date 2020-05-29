package sky

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/retry"
)

func (s *SkyConfig)CreateInformer(){
	factory := informers.NewSharedInformerFactory(s.Client, 0)
	s.Informer = factory.Core().V1().ConfigMaps()
	s.NamespaceInformer = factory.Core().V1().Namespaces()
	go factory.Start(s.Stop)
	if !cache.WaitForCacheSync(s.Stop, s.Informer.Informer().HasSynced,s.NamespaceInformer.Informer().HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	s.Informer.Informer().AddEventHandler(s)
}

func (s *SkyConfig)CheckNs()error{
	if _,err := s.NamespaceInformer.Lister().Get("sky");err != nil {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			_,err := s.Client.CoreV1().Namespaces().Create(&v1.Namespace{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
				},
				ObjectMeta:metav1.ObjectMeta{
					Name:                       s.NameSpace,
					Namespace:                  s.NameSpace,
				},
			})
			return err
		})
		return err
	}
	return nil
}

func (s *SkyConfig)CreateUserDb()error{
	labels := make(map[string]string,1)
	labels["key"] = "user"
	if err := s.Informer.Lister().ConfigMaps("sky");err != nil {
		err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			_,err := s.Client.CoreV1().ConfigMaps("sky").Create(&v1.ConfigMap{
				TypeMeta:   metav1.TypeMeta{
					Kind:   "ConfigMap",
					APIVersion:"v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:"user",
					Namespace:"sky",
					UID:uuid.NewUUID(),
					Labels:labels,
				},
				Data: map[string]string{},
			})
			return err
		})
		return err
	}
	return nil
}