# Integer
`go run sorter.go -int 1 0 3 4 7`

Output: `0 1 3 4 7`

# Float
`go run sorter.go -float 1.1 0.5 3.2 4.4 7.5`

Output: `0.5 1.1 3.2 4.4 7.5`


# String
`go run sorter.go -string apple dog cat bird egg`

Output: `apple bird cat dog egg`


# Mix
`go run sorter.go mix 1 apple 1.1 banana 0.5 cucumber 3.2 dog 4.4 7.5 5`

Output: `0.5 1 1.1 3.2 4.4 5 7.5 apple banana cucumber`

# Error: Invalid data type

`go run sorter.go -int 1 0 3 4 7 hehe`

Output: `Invalid datatype - exit status 1`