# How It Came About

Although receiving web hooks were typically a pain point in systems I have work on in the past, I never had the space to solve this problem properly.

[../images/simple_view.png]

Doesn't seem that complicated (famous last words).

I already implemented it in the easiest form I could thing of. It feels like cheating connecting libraries with half a dozen function of glue code.

Now the hard part. I will build a testing framework around this to get real working data to guide the optimization strategies to explore.

 | lib     | now     | plan       |
 | ------- | ------- | ---------- |
 | Router  | Chi     | ?          |
 | Storage | sqlite3 | ClickHouse |
 | Stream  | nats    | Kafka?     |


