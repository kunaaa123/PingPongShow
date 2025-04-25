# 📊 Ping Show - เกมจำลองปิงปองด้วย Go

เกมจำลองการแข่งขันปิงปองระหว่างผู้เล่น 2 คน พัฒนาโดยใช้ภาษา Go  
สามารถรันได้ทั้งแบบรวมและแยก service พร้อมบันทึกผลการแข่งขันลงไฟล์ CSV

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
2️⃣ ติดตั้ง Dependencies
bash
คัดลอก
แก้ไข
go mod tidy
3️⃣ ตั้งค่าฐานข้อมูล
โหลดไฟล์ database และ import เข้า MySQL

🎮 การรันเกม
วิธีที่ 1: รันแบบ All-in-One
bash
คัดลอก
แก้ไข
go run cmd/all_in_one/main.go
วิธีที่ 2: รันแต่ละส่วนแยกกัน
รัน Player Service:

bash
คัดลอก
แก้ไข
go run cmd/player_service/main.go
รัน Table Service:

bash
คัดลอก
แก้ไข
go run cmd/table_service/main.go
รัน Client:

bash
คัดลอก
แก้ไข
go run cmd/client/main.go
📖 กฎการเล่น
เกมจำลองการแข่งขันปิงปองระหว่าง 2 ผู้เล่น

ผู้เล่น A จะเริ่มด้วยการ "ping" พร้อมค่าพลังงานแบบสุ่ม

โต๊ะจะตอบกลับด้วยค่าพลังงานที่ลดลง

ผู้เล่น B จะตอบกลับด้วย "pong" พร้อมค่าพลังงานสุ่ม

เกมจะดำเนินไปจนกว่าผู้เล่นคนใดคนหนึ่งจะพ่ายแพ้ (พลังงานต่ำเกินกำหนด)

📁 โครงสร้างโปรเจค
csharp
คัดลอก
แก้ไข
pingshow/
├── cmd/              # จุดเริ่มต้นของแต่ละ service
│   ├── all_in_one/
│   ├── player_service/
│   ├── table_service/
│   └── client/
├── internal/         # โค้ดหลักของแอปพลิเคชัน
├── pkg/              # แพ็คเกจที่ใช้ร่วมกัน
├── database/         # โครงสร้างและสคริปต์ฐานข้อมูล
└── match_log.csv     # ไฟล์บันทึกผลการแข่งขัน
🗂️ การบันทึกข้อมูล
ข้อมูลการแข่งขันจะถูกบันทึกลงในไฟล์ match_log.csv ซึ่งประกอบไปด้วย:

เวลา (Timestamp)

เหตุการณ์ (Event: start_game, ping, pong, lose, game_over)

ผู้เล่น (Player)

ค่าพลังงาน (Energy)

Goroutine ID

หมายเลขการแข่งขัน (Match No.)

รอบการเล่น (Round)

---