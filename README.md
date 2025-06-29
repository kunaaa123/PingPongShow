# 📊 Ping Show - เกมจำลองปิงปองด้วย Go

"Ping Show" คือเกมจำลองการแข่งขันปิงปองระหว่างผู้เล่น 2 คน พัฒนาโดยใช้ภาษา Go 🎮  
รองรับทั้งการรันแบบรวม (All-in-One) และแยก Service พร้อมความสามารถในการบันทึกผลการแข่งขันลงไฟล์ CSV 📁

---

## 📝 ข้อกำหนดเบื้องต้น

- **Go** เวอร์ชัน `1.16` หรือใหม่กว่า  
- **MySQL** หรือฐานข้อมูลอื่นที่รองรับ  

---

## 🚀 การติดตั้ง

### 1️⃣ โคลนโปรเจค

```bash
git clone https://github.com/kunaaa123/PingPongShow
cd pingshow

```

2️⃣ ติดตั้ง Dependencies
```bash
go mod tidy

```

3️⃣ ตั้งค่าฐานข้อมูล
```bash
- โหลดไฟล์ Database
- Import เข้า MySQL
```


### วิธีที่ 1: รันแบบ All-in-One  
```bash
go run cmd/all_in_one/main.go

```
### วิธีที่ 2: รันแต่ละ Service แยกกัน
- รัน Player Service:go run cmd/player_service/main.go

```bash

- รัน Table Service:go run cmd/table_service/main.go

```
```bash
- รัน Client:go run cmd/client/main.go
```


📖 กฎการเล่น
- เกมจะเริ่มต้นเมื่อผู้เล่น A ส่ง "ping" พร้อมค่าพลังงานแบบสุ่ม
- โต๊ะจะตอบกลับด้วยค่าพลังงานที่ลดลง
- ผู้เล่น B จะส่ง "pong" พร้อมค่าพลังงานแบบสุ่ม
- เกมจะดำเนินไปจนกว่าผู้เล่นคนใดคนหนึ่งจะพ่ายแพ้ (พลังงานต่ำเกินกำหนด)



🗂️ การบันทึกข้อมูล
ข้อมูลการแข่งขันจะถูกบันทึกลงในไฟล์ match_log.csv ซึ่งประกอบไปด้วย:
- เวลา (Timestamp)
- เหตุการณ์ (Event: start_game, ping, pong, lose, game_over)
- ผู้เล่น (Player)
- ค่าพลังงาน (Energy)
- Goroutine ID
- หมายเลขการแข่งขัน (Match No.)
- รอบการเล่น (Round)

---

//test
///tttt
