package authentification

import "crypto/rand"


func GenerateKey(size int) ([]byte, error) {
    key := make([]byte, size)
    _, err := rand.Read(key)
    if err != nil {
        return nil, err // Don't forget to handle errors properly!
    }
    return key, nil
}
