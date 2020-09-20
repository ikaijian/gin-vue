package service

import (
	"fmt"
	"github.com/kaijian/gin-vue/model"
	"github.com/kaijian/gin-vue/utils"
	"sync"
)

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint, error) {
	infos := make([]*model.UserInfo, 0)

	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}
	fmt.Println(ids)
	wg := sync.WaitGroup{}

	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			shortId, err := utils.GenShortId()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()

			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  username,
				SayHello:  fmt.Sprintf("hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2020-09-19 12:00:00"),
				UpdatedAt: u.UpdatedAt.Format("2020-09-19 12:00:00"),
			}
		}(u)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil
}
