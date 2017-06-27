package watch

import (
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

var (
	PvcCache map[string]v1.PersistentVolumeClaim
)

func Pvcs(clientset *kubernetes.Clientset) chan struct{} {
	PvcCache = map[string]v1.PersistentVolumeClaim{}

	watchlist := cache.NewListWatchFromClient(clientset.Core().RESTClient(), "persistentvolumeclaims", v1.NamespaceDefault, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.PersistentVolumeClaim{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: pvcFilterAddDelete(func(pvc v1.PersistentVolumeClaim) {
				PvcCache[pvc.Name] = pvc
			}),
			DeleteFunc: pvcFilterAddDelete(func(pvc v1.PersistentVolumeClaim) {
				delete(PvcCache, pvc.Name)
			}),
			UpdateFunc: pvcFilterUpdate(func(pvc v1.PersistentVolumeClaim) {
				PvcCache[pvc.Name] = pvc
			}),
		},
	)

	stop := make(chan struct{})
	go controller.Run(stop)

	return stop
}

func pvcFilterAddDelete(f func(v1.PersistentVolumeClaim)) func(interface{}) {
	return func(obj interface{}) {
		pvc := obj.(*v1.PersistentVolumeClaim)
		f(*pvc)
	}
}

func pvcFilterUpdate(f func(v1.PersistentVolumeClaim)) func(interface{}, interface{}) {
	return func(oldObj, newObj interface{}) {
		pvcFilterAddDelete(f)(newObj)
	}
}