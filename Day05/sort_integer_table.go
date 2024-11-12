package sprint

func SortIntegerTable(table []int) []int {
    sortedTable := make([]int, len(table))
    copy(sortedTable, table)

    n := len(sortedTable)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if sortedTable[j] > sortedTable[j+1] {
                sortedTable[j], sortedTable[j+1] = sortedTable[j+1], sortedTable[j]
            }
        }
    }

    return sortedTable 
}
