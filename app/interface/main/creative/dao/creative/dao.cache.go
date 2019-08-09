// Code generated by $GOPATH/src/go-common/app/tool/cache/gen. DO NOT EDIT.

/*
  Package creative is a generated cache proxy package.
  It is generated from:
  type _cache interface {
		// cache: -sync=true -nullcache=&BgmData{Data:""} -check_null_code=$.Data==""
		BgmData(c context.Context, aid, cid, mtype int64) (*BgmData, error)
	}
*/

package creative

import (
	"context"

	"go-common/library/log"
	"go-common/library/stat/prom"
)

var _ _cache

// BgmData get data from cache if miss will call source method, then add to cache.
func (d *Dao) BgmData(c context.Context, aid, cid, mtype int64, cache bool) (res *BgmData, err error) {
	if !cache {
		res, err = d.RawBgmData(c, aid, cid, mtype)
		return
	}
	addCache := true
	res, err = d.CacheBgmData(c, aid, cid, mtype)
	if err != nil {
		addCache = false
		err = nil
	}
	defer func() {
		if res.Data == "" {
			res = nil
		}
	}()
	if res != nil {
		log.Info("aid(%d) BgmData from cache (%+v)", aid, res)
		prom.CacheHit.Incr("BgmData")
		return
	}
	prom.CacheMiss.Incr("BgmData")
	res, err = d.RawBgmData(c, aid, cid, mtype)
	if err != nil {
		log.Info("aid(%d) BgmData  in db  (%+v)", aid, res)
		return
	}
	miss := res
	if miss == nil {
		miss = &BgmData{Data: ""}
	}
	if !addCache {
		return
	}
	log.Info("aid(%d) BgmData write in cache  (%+v)", aid, miss)
	d.AddCacheBgmData(c, aid, cid, mtype, miss)
	return
}
