Transaction logic:

1.创世区块
transactions
{
    Id: 00000
    Vin [TXInput{[]byte{}, -1, time.Now()}]
    Vout [TXOutput{10, wangkun}]
}

2.Block
transactions
{
    Id: 11111
    Vin [TXInput{00000, 1, wangkun}]
    Vout [TXOutPut{6, min}, TXOupt{4, wangkun}]
}

结论：一个TXOutPut只能花费一次，换句话说一个TXOutPut如果要使用它，那么就建一个与之对应的TXInput

3.Block
transactions
{
    Id: 22222
    Vin [TXInput{1111, 1, min}]
    Vout [TXOutPut{1, kun}, TXOutput{5, min}]
}


...

查询：
1.查询wnagkun账号里面包含未花费的所有transaction
2.需要判断哪一个TXOutput



