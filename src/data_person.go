package main

type person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends people `json:"friends"`
}

func (p person) getFriends() (people, error) {
	list := people{}

	if refs, found := dataFriendship[p.ID]; found {
		for _, friendID := range refs {
			friend, err := dataPeople.byID(friendID)

			if err != nil {
				return nil, err
			}

			list = append(list, friend)

		}
	}

	return list, nil
}
