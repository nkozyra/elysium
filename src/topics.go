package elysium

import (
	"html/template"
	"log"
)

type Topic struct {
	ID        int     `json:"topic_id"`
	Title     string  `json:"topic_title"`
	Guid      string  `json:"topic_guid"`
	ForumID   int     `json:"forum_id"`
	ForumGuid string  `json:"forum_guid"`
	ForumName string  `json:"forum_name"`
	User      int     `json:"user_id"`
	Time      int     `json:"topic_create_time"`
	Topics    []Topic `json:"topic_topics"`
	Posts     []Post  `json:"posts"`
}

func (t Topic) create(user_id int, title string, description string) {
	t.Title = title
	t.Guid = t.GenerateGUID()
	res, err := DB.Exec("INSERT INTO topics SET user_id=?, forum_id=1, topic_title=?, topic_guid=?", user_id, title, t.Guid)
	if err != nil {
		log.Println(err)
	}
	topic_id, _ := res.LastInsertId()
	_, err = DB.Exec("INSERT INTO posts SET user_id=?, topic_id=?, post_title=?, post_text=?", user_id, topic_id, title, description)
	if err != nil {
		log.Println(err)
	}
}

func (t Topic) GenerateGUID() string {
	sg := simpleGUID(t.Title)
	rows, err := DB.Query("SELECT topic_guid FROM topics WHERE topic_guid=?", sg)
	if err != nil {
		// handle
	}
	rowCount := 0
	for rows.Next() {
		rowCount++
	}
	if rowCount > 0 {
		return sg + RandStringBytes(6)
	} else {
		return sg
	}
}

func (t Topic) GetPosts() []Post {

	rows, err := DB.Query("SELECT post_id, post_text, user_id FROM posts WHERE topic_id=? ORDER BY post_create_time DESC", t.ID)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		p := Post{}
		rows.Scan(&p.ID, &p.Text, &p.User)
		p.Text = nl2br(p.Text)
		p.Text = parseText(p.Text)
		p.FText = template.HTML(p.Text)
		t.Posts = append(t.Posts, p)
	}
	return t.Posts
}

func GetTopic(guid string) Topic {
	t := Topic{}
	err := DB.QueryRow("SELECT topic_id, topic_guid, topic_title, t.forum_id, f.forum_guid, f.forum_name FROM topics t LEFT JOIN forums f ON f.forum_id=t.forum_id WHERE topic_guid=?", guid).Scan(&t.ID, &t.Guid, &t.Title, &t.ForumID, &t.ForumGuid, &t.ForumName)
	if err != nil {
		log.Println(err)
	}

	t.Posts = t.GetPosts()
	f := GetForum(t.ForumGuid, 1)
	t.Topics = f.Topics

	return t
}
