package authentification

import("regexp")

func validatePassword(password string) (bool, string) {
    var (
        uppercase = regexp.MustCompile(`[A-Z]`)
        lowercase = regexp.MustCompile(`[a-z]`)
        special   = regexp.MustCompile(`[!@#\$%\^&*()_+{}\[\]:;<>,.?/\|~-]`)
    )

    if !uppercase.MatchString(password) {
        return false, "Le mot de passe doit contenir au moins une majuscule."
    }

    if !lowercase.MatchString(password) {
        return false, "Le mot de passe doit contenir au moins une minuscule."
    }

    if !special.MatchString(password) {
        return false, "Le mot de passe doit contenir au moins un caractère spécial."
    }

    // Le mot de passe passe toutes les vérifications
    return true, ""
}

