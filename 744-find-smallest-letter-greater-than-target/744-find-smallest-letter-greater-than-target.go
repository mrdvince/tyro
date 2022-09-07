func nextGreatestLetter(letters []byte, target byte) byte {
    for _, letter := range letters {
        if target < letters[0] {
            break
        }
        if letter > target {
            return letter
        }
    }
    return letters[0]
}