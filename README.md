# Package

Package ในภาษา Go ใช้เพื่อจัดกลุ่ม source code ที่ทำงานเรื่องเดียวกัน
ให้มารวมอยู่ใน directory เดียวกัน ทำให้ง่ายต่อการอ่าน และทำความเข้าใจ

## import

การ import package แบ่งออกเป็น 3 รูปแบบ ของ package

1. built-in ที่มาพร้อมกับภาษา Go
2. package ที่เราแบ่ง และสร้างขึ้น เองใน project
3. package ที่เป็น open-source

### 1. build-in package

ยกตัวอย่างเช่น

```go
import "fmt"
```

เนื่องจากเป็น built-in package เราจึงสามารถอ้างเพียงแค่ชื่อได้เลย

### 2. project package

เป็น package ที่เราเขียนและสร้างขึ้นภายใน project ของเราเอง เช่น book
โดยวิธีการแยก package มีขั้นตอนคือ

- สร้าง directory (ซึ่งต้องอยู่ภายใต้ directory ที่มีไฟล์ go.mod อยู่)
- สร้าง file ที่มี extension นามสกุล .go
- บรรทัดแรกของไฟล์ .go จะต้องประกาศชื่อ package เข่น `package book`
- *ชื่อ package ควรตั้งให้ตรงกับชื่อของ directory*
- ข้อแนะนำการตั้งชื่อคือ ใช้อักษรตัวเล็กในภาษาอังกฤษทั้งหมดแบบไม่มีการคั่นคำ เช่น `printableobject` เป็นต้น
- วิธีการ import ให้เอาชื่อของ module(ดูที่ไฟล์ go.mod) มาตั้งต้นแล้วต่อด้วย relative path ไปที่ directory ของ package เช่น 

```go
 import "github.com/pallat/skooldio/packagedemo/book"
```

### 3. open-source package

เป็นการ import จาก open-source ต่างๆเข้ามาใช้ เช่น

```go
import "github.com/gin-gonic/gin"
```

โดยการอ้างถึง open-source จะต้องอ้างถึงที่อยู่ของ source code นั้นแบบ URL
