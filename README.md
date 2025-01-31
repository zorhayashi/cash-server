## cash-server system 
#### GO (Gin) 
##### 資料夾
- routers 路由表
- configs 設定檔
- controller 
- database 資料庫
- doc swagger（*不要動）
- model  
- services
- grpc (casino連線)
- logs (紀錄)
- vendor 外部套件（*不要動）
- templates web模組
- pkg 套件
    - e  錯誤訊息（未實裝）
    - encryption  加密演算
    - line  通知
    - flag  
    - util  
      - log
      - name
      - time
      - logger
- - - 

### 註冊

[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdHJlY3QgcmdiKDIxMCwgMjU1LCAyNDApXG5cdE5vdGUgcmlnaHQgb2YgV2ViX1NlcnZlcjogUmVnaXN0ZXJlZCAoL2FkbWluL3JlZ2lzdGVyKVxuXHRXZWJfU2VydmVyLT4-K0Nhc2hfU2VydmVyOiBuYW1lICwgYWNjb3VudCAsIHBhc3N3b3JkXG5cdENhc2hfU2VydmVyLS0-PitXZWJfU2VydmVyOiBUb2tlbiAsIFRva2VuU2VjcmV0XG5cdGVuZFxuXG5cblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImZvcmVzdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cdHJlY3QgcmdiKDIxMCwgMjU1LCAyNDApXG5cdE5vdGUgcmlnaHQgb2YgV2ViX1NlcnZlcjogUmVnaXN0ZXJlZCAoL2FkbWluL3JlZ2lzdGVyKVxuXHRXZWJfU2VydmVyLT4-K0Nhc2hfU2VydmVyOiBuYW1lICwgYWNjb3VudCAsIHBhc3N3b3JkXG5cdENhc2hfU2VydmVyLS0-PitXZWJfU2VydmVyOiBUb2tlbiAsIFRva2VuU2VjcmV0XG5cdGVuZFxuXG5cblxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImZvcmVzdCJ9LCJ1cGRhdGVFZGl0b3IiOmZhbHNlfQ)

### 要求商品清單

[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cblx0cmVjdCByZ2IoMjEwLCAyNTUsIDI0MClcblx0Tm90ZSByaWdodCBvZiBXZWJfU2VydmVyOiBHZXQgSXRlbSBsaXN0IFxuXHRXZWJfU2VydmVyLT4-K0Nhc2hfU2VydmVyOiBUb2tlbiBcblx0XHRcdHJlY3QgcmdiKDIxLCAyNTUsIDI0MClcblx0XHRcdFx0XHROb3RlIHJpZ2h0IG9mIENhc2hfU2VydmVyOiBHUlBDXG5cdFx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K0dhbWVfU2VydmVyOiAgR2V0IEl0ZW0gbGlzdFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBsaXN0XG5cdFx0XHRlbmRcblx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K1dlYl9TZXJ2ZXI6IHJldHVybiBsaXN0XG5cdGVuZFxuXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJmb3Jlc3QifSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cblx0cmVjdCByZ2IoMjEwLCAyNTUsIDI0MClcblx0Tm90ZSByaWdodCBvZiBXZWJfU2VydmVyOiBHZXQgSXRlbSBsaXN0IFxuXHRXZWJfU2VydmVyLT4-K0Nhc2hfU2VydmVyOiBUb2tlbiBcblx0XHRcdHJlY3QgcmdiKDIxLCAyNTUsIDI0MClcblx0XHRcdFx0XHROb3RlIHJpZ2h0IG9mIENhc2hfU2VydmVyOiBHUlBDXG5cdFx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K0dhbWVfU2VydmVyOiAgR2V0IEl0ZW0gbGlzdFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBsaXN0XG5cdFx0XHRlbmRcblx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K1dlYl9TZXJ2ZXI6IHJldHVybiBsaXN0XG5cdGVuZFxuXG5cbiIsIm1lcm1haWQiOnsidGhlbWUiOiJmb3Jlc3QifSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)

###  order

[![](https://mermaid.ink/img/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cblx0cmVjdCByZ2IoMjEwLCAyNTUsIDI0MClcblx0Tm90ZSByaWdodCBvZiBXZWJfU2VydmVyOiBidXkgb3JkZXIoL015Y2FyZC8xMjMpXG5cdFdlYl9TZXJ2ZXItPj4rQ2FzaF9TZXJ2ZXI6IFRva2VuICwgSXRlbSBJRCAsIFVzZXJJRFxuXHRcdFx0XHRyZWN0IHJnYigyMzAsIDIzNSwgMjEwKVxuXHRcdFx0XHRcdE5vdGUgcmlnaHQgb2YgQ2FzaF9TZXJ2ZXI6IOeiuuiqjeW4s-iZn-WtmOWcqFxuXHRcdFx0XHRcdENhc2hfU2VydmVyLS0-PitHYW1lX1NlcnZlcjogIGNoZWNrIFVzZXJJRFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBUL0Zcblx0XHRcdGVuZFxuXHRcdFx0XHRDYXNoX1NlcnZlci0tPj4rV2ViX1NlcnZlcjogPEZhbHNlPiByZXR1cm4gTm9cblx0XHRcdFx0Q2FzaF9TZXJ2ZXItPj4rT3RoZW5QYXlfU2VydmVyOiAgcG9zdCAgMy4xXG5cdFx0XHRcdE90aGVuUGF5X1NlcnZlci0tPj4rQ2FzaF9TZXJ2ZXI6ICBBdXRoQ29kZSBcblx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K1dlYl9TZXJ2ZXI6ICBSZWRpcmVjdF91cmwgKyBBdXRoQ29kZSBcblx0XHRcdFx0V2ViX1NlcnZlci0-PitPdGhlblBheV9TZXJ2ZXI6ICBHZXQgV2ViIDMuMlxuXHRcdFx0XHRPdGhlblBheV9TZXJ2ZXItLT4-K090aGVuUGF5X1NlcnZlcjogIHdlYiDmtojosrvkuqTmmJNcblx0XHRcdFx0T3RoZW5QYXlfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogIDMuMSAgRmFjUmV0dXJuVVJMXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitPdGhlblBheV9TZXJ2ZXI6ICBwb3N0IEF1dGhDb2RlICAzLjMgXG5cdFx0XHRcdE90aGVuUGF5X1NlcnZlci0tPj4rQ2FzaF9TZXJ2ZXI6ICBtc2cgXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitPdGhlblBheV9TZXJ2ZXI6ICBwb3N0IEF1dGhDb2RlICAzLjRcblx0XHRcdFx0T3RoZW5QYXlfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogIG1zZyBcblx0XHRcdFx0cmVjdCByZ2IoMjMwLCAyMzUsIDIxMClcblx0XHRcdFx0XHROb3RlIHJpZ2h0IG9mIENhc2hfU2VydmVyOiDnmbzpgIHllYblk4Hkv6Hku7Zcblx0XHRcdFx0XHRDYXNoX1NlcnZlci0tPj4rR2FtZV9TZXJ2ZXI6ICBQT1NUIFVzZXJJRFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBUL0Zcblx0XHRcdFx0ZW5kXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitXZWJfU2VydmVyOiByZXR1cm4gaXNPS1xuXHRlbmRcbiIsIm1lcm1haWQiOnsidGhlbWUiOiJmb3Jlc3QifSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)](https://mermaid-js.github.io/mermaid-live-editor/#/edit/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG5cblx0cmVjdCByZ2IoMjEwLCAyNTUsIDI0MClcblx0Tm90ZSByaWdodCBvZiBXZWJfU2VydmVyOiBidXkgb3JkZXIoL015Y2FyZC8xMjMpXG5cdFdlYl9TZXJ2ZXItPj4rQ2FzaF9TZXJ2ZXI6IFRva2VuICwgSXRlbSBJRCAsIFVzZXJJRFxuXHRcdFx0XHRyZWN0IHJnYigyMzAsIDIzNSwgMjEwKVxuXHRcdFx0XHRcdE5vdGUgcmlnaHQgb2YgQ2FzaF9TZXJ2ZXI6IOeiuuiqjeW4s-iZn-WtmOWcqFxuXHRcdFx0XHRcdENhc2hfU2VydmVyLS0-PitHYW1lX1NlcnZlcjogIGNoZWNrIFVzZXJJRFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBUL0Zcblx0XHRcdGVuZFxuXHRcdFx0XHRDYXNoX1NlcnZlci0tPj4rV2ViX1NlcnZlcjogPEZhbHNlPiByZXR1cm4gTm9cblx0XHRcdFx0Q2FzaF9TZXJ2ZXItPj4rT3RoZW5QYXlfU2VydmVyOiAgcG9zdCAgMy4xXG5cdFx0XHRcdE90aGVuUGF5X1NlcnZlci0tPj4rQ2FzaF9TZXJ2ZXI6ICBBdXRoQ29kZSBcblx0XHRcdFx0Q2FzaF9TZXJ2ZXItLT4-K1dlYl9TZXJ2ZXI6ICBSZWRpcmVjdF91cmwgKyBBdXRoQ29kZSBcblx0XHRcdFx0V2ViX1NlcnZlci0-PitPdGhlblBheV9TZXJ2ZXI6ICBHZXQgV2ViIDMuMlxuXHRcdFx0XHRPdGhlblBheV9TZXJ2ZXItLT4-K090aGVuUGF5X1NlcnZlcjogIHdlYiDmtojosrvkuqTmmJNcblx0XHRcdFx0T3RoZW5QYXlfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogIDMuMSAgRmFjUmV0dXJuVVJMXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitPdGhlblBheV9TZXJ2ZXI6ICBwb3N0IEF1dGhDb2RlICAzLjMgXG5cdFx0XHRcdE90aGVuUGF5X1NlcnZlci0tPj4rQ2FzaF9TZXJ2ZXI6ICBtc2cgXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitPdGhlblBheV9TZXJ2ZXI6ICBwb3N0IEF1dGhDb2RlICAzLjRcblx0XHRcdFx0T3RoZW5QYXlfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogIG1zZyBcblx0XHRcdFx0cmVjdCByZ2IoMjMwLCAyMzUsIDIxMClcblx0XHRcdFx0XHROb3RlIHJpZ2h0IG9mIENhc2hfU2VydmVyOiDnmbzpgIHllYblk4Hkv6Hku7Zcblx0XHRcdFx0XHRDYXNoX1NlcnZlci0tPj4rR2FtZV9TZXJ2ZXI6ICBQT1NUIFVzZXJJRFxuXHRcdFx0XHRcdEdhbWVfU2VydmVyLS0-PitDYXNoX1NlcnZlcjogUmVzcCBUL0Zcblx0XHRcdFx0ZW5kXG5cdFx0XHRcdENhc2hfU2VydmVyLS0-PitXZWJfU2VydmVyOiByZXR1cm4gaXNPS1xuXHRlbmRcbiIsIm1lcm1haWQiOnsidGhlbWUiOiJmb3Jlc3QifSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)
