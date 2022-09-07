func nextGreatestLetter(letters []byte, target byte) byte {
    point := sort.Search(len(letters), func(i int) bool { return letters[i] > target })
    if point == len(letters) {
        return letters[0]
    } else {
        return letters[point]
    }
}