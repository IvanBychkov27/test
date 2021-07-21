package main

//
//type ScoreOS struct {
//	Stamp               int // Combined.Ts - индификатор для поиска в базе
//	Score               int // Рейтинг от 0 до 100
//	PlatformName        string
//	PlatformVersion     string
//	PlatformVersionName string
//}
//
//// расчет рейтинга
//func fingerprint(uabase []Combined, user Combined) []ScoreOS {
//	scores := make([]ScoreOS, 0, len(uabase))
//	for _, ua := range uabase {
//		score := 0
//		if ua.Os.Name == "" {
//			continue
//		}
//
//		if ua.IPDf == user.IPDf {
//			score += 10
//		}
//		if ua.IPMf == user.IPMf {
//			score += 10
//		}
//		if ua.TCPWindowSize == user.TCPWindowSize {
//			score += 15
//		}
//		if ua.TCPFlags == user.TCPFlags {
//			score += 10
//		}
//		if ua.TCPHeaderLength == user.TCPHeaderLength {
//			score += 10
//		}
//		if ua.TCPMss == user.TCPMss {
//			score += 15
//		}
//		if ua.TCPOptions == user.TCPOptions {
//			score += 30
//		}
//
//		// проверяем порядок параметров TCP (это слабее, чем равенство параметров TCP)
//		if ua.TCPOptions != user.TCPOptions && ua.TCPOptions != "" && user.TCPOptions != "" {
//			var orderUA, orderUser string
//			for _, e := range strings.Split(ua.TCPOptions, ",") {
//				if e != "" {
//					orderUA += string(e[0])
//				}
//			}
//			for _, e := range strings.Split(user.TCPOptions, ",") {
//				if e != "" {
//					orderUser += string(e[0])
//				}
//			}
//			if orderUA == orderUser {
//				score += 20
//			}
//		}
//
//		sc := ScoreOS{
//			Stamp:               ua.Ts,
//			Score:               score,
//			PlatformName:        ua.Os.Name,
//			PlatformVersion:     ua.Os.Version,
//			PlatformVersionName: ua.Os.VersionName,
//		}
//		scores = append(scores, sc)
//	}
//	return scores
//}
//
//// вывод n элементов по ТОП рейтингу - Инфо: рейтинг ОС, Имя ОС и версия
//func resultOS(scores []ScoreOS, n int) {
//	sort.SliceStable(scores, func(i, j int) bool {
//		return scores[i].Score > scores[j].Score // сортировка по убыванию рейтинга
//	})
//
//	i := 0
//	for _, sc := range scores {
//		fmt.Printf("score %3d  OS: %s %s  %s \n", sc.Score, sc.PlatformName, sc.PlatformVersionName, sc.PlatformVersion)
//		i++
//		if i == n {
//			break
//		}
//	}
//}
