package sky

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"

	"k8s.io/client-go/tools/cache"
)

func (s *SkyConfig)CreateInformer(){
	factory := informers.NewSharedInformerFactory(s.Client, 0)
	s.Informer = factory.Core().V1().ConfigMaps()
	go factory.Start(s.Stop)
	if !cache.WaitForCacheSync(s.Stop, s.Informer.Informer().HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
}

func (s *SkyConfig)AddFunc(object interface{}){

}


func (s *SkyConfig)UpdateFunc(object interface{}){

}


func (s *SkyConfig)DeleteFunc(object interface{}){

}

func (s *SkyConfig)CheckNs(){
	if _,err := s.Informer.Lister().ConfigMaps("sky").Get("");err != nil {
		s.Client.CoreV1().Namespaces().Create(&v1.Namespace{
			TypeMeta: metav1.TypeMeta{
					Kind:       "Namespace",
					APIVersion: "v1",
			},
			ObjectMeta:metav1.ObjectMeta{
				Name:                       s.NameSpace,
				Namespace:                  s.NameSpace,
			},
		})
	}

}
