# Go-reloaded

Un outil en Go qui corrige un texte. Il lit un fichier d'entrée, applique des
règles de modification et écrit le texte corrigé dans un fichier de sortie.

Le projet utilise uniquement la bibliothèque standard de Go.

## Utilisation

```bash
go run . sample.txt result.txt
```

- `sample.txt` : le fichier d'entrée, avec le texte et ses balises
- `result.txt` : le fichier de sortie, avec le texte corrigé

Exemple :

```
$ cat sample.txt
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
```

Sur Windows, on peut aussi lancer `start.bat`.

## Les règles

### Conversions de nombres

| Balise  | Effet                                   | Entrée                 | Sortie              |
|---------|-----------------------------------------|------------------------|---------------------|
| `(hex)` | Le mot d'avant passe de base 16 à base 10 | `1E (hex) files`     | `30 files`          |
| `(bin)` | Le mot d'avant passe de base 2 à base 10  | `It has been 10 (bin) years` | `It has been 2 years` |

### Majuscules et minuscules

| Balise  | Effet                         | Entrée                           | Sortie                          |
|---------|-------------------------------|----------------------------------|---------------------------------|
| `(up)`  | Le mot d'avant en majuscules  | `Ready, set, go (up) !`          | `Ready, set, GO!`               |
| `(low)` | Le mot d'avant en minuscules  | `I should stop SHOUTING (low)`   | `I should stop shouting`        |
| `(cap)` | Majuscule au début du mot     | `Welcome to the Brooklyn bridge (cap)` | `Welcome to the Brooklyn Bridge` |

On peut ajouter un nombre pour modifier plusieurs mots :
`(up, 2)`, `(low, 3)`, `(cap, 6)`...

```
This is so exciting (up, 2)  ->  This is SO EXCITING
```

### Ponctuation

Les signes `. , ! ? : ;` sont collés au mot d'avant et suivis d'un espace.

```
over there ,and then BAMM !!  ->  over there, and then BAMM!!
```

Les groupes comme `...` ou `!?` restent entiers.

```
I was thinking ... You were right  ->  I was thinking... You were right
```

### Apostrophes

Les apostrophes vont par paire et se collent aux mots qu'elles entourent.

```
describe me: ' awesome '  ->  describe me: 'awesome'
```

### « a » devient « an »

`a` devient `an` quand le mot suivant commence par une voyelle (`a e i o u`) ou un `h`.

```
There it was. A amazing rock!  ->  There it was. An amazing rock!
```

### Erreur

Une balise avec un nombre négatif, comme `(up, -1)`, n'est pas valide.
La ligne est alors écrite entre crochets, suivie de `=err syntax`.

## Organisation du code

Chaque fonction est dans son propre fichier, qui porte son nom.

| Fichier                     | Rôle                                                  |
|-----------------------------|-------------------------------------------------------|
| `main.go`                   | Vérifie les arguments, lit et écrit les fichiers      |
| `transformText.go`          | Corrige le texte ligne par ligne                      |
| `transformLine.go`          | Applique toutes les étapes sur une ligne              |
| `separatePunctuation.go`    | `,and` devient `,` et `and`                           |
| `joinTagWords.go`           | `(up,` et `2)` deviennent `(up,2)`                    |
| `separateTagPunctuation.go` | `(up),` devient `(up)` et `,`                         |
| `hasNegativeTag.go`         | Détecte une balise avec un nombre négatif             |
| `getTagNumber.go`           | Lit le nombre d'une balise comme `(up,2)`             |
| `applyTags.go`              | Applique les balises                                  |
| `hexToDecimal.go`           | Conversion hexadécimal vers décimal                   |
| `binToDecimal.go`           | Conversion binaire vers décimal                       |
| `toUpperWords.go`           | Met des mots en majuscules                            |
| `toLowerWords.go`           | Met des mots en minuscules                            |
| `capitalizeWords.go`        | Met une majuscule au début de plusieurs mots          |
| `capitalizeWord.go`         | Met une majuscule au début d'un mot                   |
| `fixArticles.go`            | Remplace `a` par `an`                                 |
| `startsWithVowelOrH.go`     | Vérifie si un mot commence par une voyelle ou un `h`  |
| `isUpperCase.go`            | Vérifie si un mot est tout en majuscules              |
| `isPunctuation.go`          | Vérifie si un mot n'est que de la ponctuation         |
| `fixPunctuation.go`         | Colle la ponctuation au mot d'avant                   |
| `fixQuotes.go`              | Colle les apostrophes aux mots                        |

## Tests

```bash
go test ./...
```

Les tests sont dans `main_test.go`.
