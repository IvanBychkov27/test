package main

//_ "github.com/lib/pq"

//
//// параметры подключения к базе данных
//func params() string {
//	pgConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
//		"dbuser",    // PgLogin
//		"secret",    // PgPassword
//		"127.0.0.1", // PgAddress
//		5432,        // PgPort
//		"test",      // PgDatabaseName
//		"disable",   // PgSSLMode
//		"",          // PgSSLCertPath
//	)
//	return pgConnString
//}
//
//// добавление всех записей из слайса в таблицу
//func addAllRecords(data []*combined.Combined) {
//	db, err := sql.Open("postgres", params()) // подключение к базе данных
//	if err == nil {
//		err = db.Ping() // отправляем запрос к базе данных - проверка подключения к базе данных
//		//fmt.Println("connecting to postgres...")
//	}
//	chk(err)
//	defer func() {
//		err = db.Close() // закрытие базы данных
//		chk(err)
//	}()
//
//	if len(data) == 0 {
//		fmt.Println("новые данные для внесения в таблицу отсутствуют")
//		return
//	}
//	var recordsCount int64
//	for _, d := range data {
//		res, err := insert(db,
//			d.Ts,
//			d.IPTTL,
//			d.IPDf,
//			d.IPMf,
//			d.TCPWindowSize,
//			d.TCPFlags,
//			d.TCPAck,
//			d.TCPSeq,
//			d.TCPHeaderLength,
//			d.TCPUrp,
//			d.TCPOptions,
//			d.TCPWindowScaling,
//			d.TCPTimestamp,
//			d.TCPTimestampEchoReply,
//			d.TCPMss,
//			d.NavigatorUserAgent,
//			d.Platform.Type,
//			d.Platform.Vendor,
//			d.Platform.Model,
//			d.DeviceMemory,
//			d.HardwareConcurrency,
//			d.Os.Name,
//			d.Os.Version,
//			d.Os.VersionName,
//		)
//
//		chk(err)
//		recordsCount += res
//	}
//	fmt.Printf("Добавлено %d записей в таблицу combineds\n", recordsCount)
//}
//
//// создает новую запись и возвращает кол-во затронутых строк и ошибку
//func insert(db *sql.DB,
//	r1, r2, r3, r4, r5, r6, r7, r8, r9, r10 int,
//	r11 string,
//	r12, r13, r14, r15 int,
//	r16, r17, r18, r19 string,
//	r20 float32,
//	r21 int,
//	r22, r23, r24 string,
//) (int64, error) {
//
//	query := `
//INSERT INTO combineds VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
//                              $11,
//                              $12, $13, $14, $15,
//                              $16, $17, $18, $19,
//                              $20,
//                              $21,
//                              $22, $23, $24)
//`
//
//	res, err := db.Exec(query,
//		r1, r2, r3, r4, r5, r6, r7, r8, r9, r10,
//		r11,
//		r12, r13, r14, r15,
//		r16, r17, r18, r19,
//		r20,
//		r21,
//		r22, r23, r24,
//		)
//	if err != nil {
//		return 0, err
//	}
//	return res.RowsAffected()
//}
//
//// обработка ошибки
//func chk(err error) {
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
