package businessrealiz

import (
	"forum/internal/models"
	"time"
)

func (s *Service) GetPostsForHome(pageNum, onPage int, categories []string) (*[]models.Post, error) {
	time.Now()
	posts := &[]models.Post{
		{
			Post_ID:      1,
			User_ID:      1,
			Post_Title:   "Violet",
			Post_Content: "!@#$%^&*()",
			Created_Time: time.Now(),
		}, {
			Post_ID:      2,
			User_ID:      2,
			Post_Title:   "Red",
			Post_Content: "-1.00",
			Created_Time: time.Now(),
		}, {
			Post_ID:      3,
			User_ID:      3,
			Post_Title:   "Yellow",
			Post_Content: "ÅÍÎÏ˝ÓÔÒÚÆ☃",
			Created_Time: time.Now(),
		}, {
			Post_ID:      4,
			User_ID:      4,
			Post_Title:   "Crimson",
			Post_Content: "åß∂ƒ©˙∆˚¬…æ",
			Created_Time: time.Now(),
		}, {
			Post_ID:      5,
			User_ID:      5,
			Post_Title:   "Fuscia",
			Post_Content: "Ṱ̺̺̕o͞ ̷i̲̬͇̪͙n̝̗͕v̟̜̘̦͟o̶̙̰̠kè͚̮̺̪̹̱̤ ̖t̝͕̳̣̻̪͞h̼͓̲̦̳̘̲e͇̣̰̦̬͎ ̢̼̻̱̘h͚͎͙̜̣̲ͅi̦̲̣̰̤v̻͍e̺̭̳̪̰-m̢iͅn̖̺̞̲̯̰d̵̼̟͙̩̼̘̳ ̞̥̱̳̭r̛̗̘e͙p͠r̼̞̻̭̗e̺̠̣͟s̘͇̳͍̝͉e͉̥̯̞̲͚̬͜ǹ̬͎͎̟̖͇̤t͍̬̤͓̼̭͘ͅi̪̱n͠g̴͉ ͏͉ͅc̬̟h͡a̫̻̯͘o̫̟̖͍̙̝͉s̗̦̲.̨̹͈̣",
			Created_Time: time.Now(),
		}, {
			Post_ID:      6,
			User_ID:      6,
			Post_Title:   "Pink",
			Post_Content: "NIL",
			Created_Time: time.Now(),
		}, {
			Post_ID:      7,
			User_ID:      7,
			Post_Title:   "Purple",
			Post_Content: "｀ｨ(´∀｀∩",
			Created_Time: time.Now(),
		}, {
			Post_ID:      8,
			User_ID:      8,
			Post_Title:   "Puce",
			Post_Content: "​",
			Created_Time: time.Now(),
		}, {
			Post_ID:      9,
			User_ID:      9,
			Post_Title:   "Pink",
			Post_Content: "⁦test⁧",
			Created_Time: time.Now(),
		}, {
			Post_ID:      10,
			User_ID:      10,
			Post_Title:   "Orange",
			Post_Content: "''",
			Created_Time: time.Now(),
		},
	}
	return posts, nil
}
