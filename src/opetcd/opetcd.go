package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.etcd.io/etcd/clientv3"
)

type Callback func(key string) string

func watchCB(key string) string {
	return key + " watch"
}

func main() {

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	kvc := clientv3.NewKV(client)

	// put
	kvc.Delete(ctx, "/user/", clientv3.WithPrefix())
	putresponse, err := kvc.Put(ctx, "/user/", "18616931990")
	defer cancel()
	if err != nil {
		log.Println("error:", err.Error())
	}
	log.Println("put revision:", putresponse.Header.Revision)

	// get
	getresponse, err := kvc.Get(ctx, "/user/")
	if err != nil {
		if err == context.Canceled {
			log.Println("error:", err.Error())
		}
		if clientv3.IsConnCanceled(err) {
			log.Println("error:", err.Error())
		}
	}

	// watch
	WatchDemo(ctx, "/user/", client, kvc, watchCB)

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend),
		clientv3.WithLimit(10),
	}

	// get key-value again
	log.Println("--------------------get range start--------------------")
	getresponses, err := kvc.Get(ctx, "/user/", opts...)
	for _, item := range getresponses.Kvs {
		if getresponses.More == true {
			log.Printf("key: %s, value:%s", string(item.Key), string(item.Value))
		}
	}
	log.Println("--------------------get range start--------------------")

	log.Println("getvsision", getresponse.Header.Revision)

	LeaseDemo(ctx, "/user/", client, kvc)

}

func WatchDemo(ctx context.Context, key string, client *clientv3.Client, kv clientv3.KV, callback Callback) {
	log.Println("Watch start...", key)
	kv.Delete(ctx, key, clientv3.WithPrefix())
	_, err := kv.Get(ctx, key)
	if err != nil {
		log.Println("get error")
	}

	//	if resp.Count == 0 {
	//		log.Println("delete ok")
	//	}

	done := make(chan int)
	go func() {
		watchChan := client.Watch(ctx, key, clientv3.WithPrefix())
		for {
			select {
			case result := <-watchChan:
				for _, ev := range result.Events {
					cbvalue := callback(string(ev.Kv.Value))
					log.Printf("%s %q : %q, cb: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value, cbvalue)
				}
			case <-done:
				log.Println("done")
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("/user/%02d", i)
		val := fmt.Sprintf("1861693199%d", i)
		kv.Put(ctx, key, val)
	}

	time.Sleep(3 * time.Second)
	done <- 1
	log.Println("Watch end")

	//Insert some more keys (no one is watching)
	for i := 10; i < 20; i++ {
		key := fmt.Sprintf("/user/%02d", i)
		val := fmt.Sprintf("186169319%02d", i)
		kv.Put(ctx, key, val)
	}

}

func LeaseDemo(ctx context.Context, key string, cli *clientv3.Client, kv clientv3.KV) {
	fmt.Println("*** LeaseDemo()")
	// Delete all keys
	kv.Delete(ctx, key, clientv3.WithPrefix())

	getresponser, _ := kv.Get(ctx, key)
	if len(getresponser.Kvs) == 0 {
		fmt.Println("No 'key'")
	}

	lease, err := cli.Grant(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Insert key with a lease of 1 second TTL
	kv.Put(ctx, key, "value"+key, clientv3.WithLease(lease.ID))

	getresponser, _ = kv.Get(ctx, key)
	if len(getresponser.Kvs) == 1 {
		fmt.Println("Found 'key'")
		log.Println("found key===", getresponser.Count)
	}

	// get
	gr, _ := kv.Get(ctx, key, clientv3.WithPrefix())
	for _, item := range gr.Kvs {
		//if gr.More == true {
		log.Printf("LeaseDemo-----key: %s, value:%s", string(item.Key), string(item.Value))
		//}
	}

	// Let the TTL expire
	time.Sleep(3 * time.Second)

	getresponser, _ = kv.Get(ctx, key)
	if len(getresponser.Kvs) == 0 {
		fmt.Println("No more 'key'")
	}
}
