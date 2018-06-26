# bucket-accounting
Software to maintain and organize "buckets" based on budgeting goals and regular spending

## Buckets, Transactions, Reports, and Accounts
### Buckets
A bucket contains a goal or other budgeting value, including a goal amount and starting amount. They also contain the name of a bank account where the appropriate money is stored, and other similar metadata. They represent the logic of your goal, and are timeless/stateless. Buckets do not care when or in what order transactions are enacted. Each bucket also has a corresponding Report. Buckets are stored on disk.

By default, a bucket is single-use. Once it fills up/hits its goal date, it won't save any more money even if it loses money in a transaction. Perhaps these single-use buckets will automatically dissolve after a transaction is made, or perhaps make it manual. Dissolved buckets are stored on disk but are inactive. Dissolved buckets can be manually reactivated; The starting date is set to the date of the transaction and the starting amount is previous "current" amount. The new goal date will be the same date plus six months. 

If a bucket is not single-use, it gets topped up. This means that they continually accept paychecks as long as their current amount is lower than the goal. Essentially, the first transaction that depletes a full bucket performs the same action of reactivation described above.

### Transactions
A transaction is any event that affects the value of a bucket. Transactions are performed on one bucket. Paychecks are special events that create positive transactions across multiple (all?) buckets. A dissolve or reactivation transaction is a special zero-sum transaction. A transaction is also stored in a bucket's report, but is not stored on disk.

### Reports
A report corresponds to a specific bucket. It keeps a record of every transaction and can be used as a ledger. The date of a report is pulled from a transaction's date, but only updates if it's later than the report's current date. When transactions are run out-of-order, the report automatically adjusts all affected totals to match the corrected timeline.

### Accounts
An account just stores the currently saved values of all buckets that reference it. It also tracks left-over funds from dissolved buckets. An account can be drained to show that leftover funds have returned to your main pool of funds. Draining an account deletes all dissolved buckets for that account. Accounts are stored on disk but are almost stateless.

### Paychecks
Will store a list of all paychecks, to be iterated over when new buckets are made (to catch them up). Any paycheck dated before the earliest bucket's starting date will be left off the list.