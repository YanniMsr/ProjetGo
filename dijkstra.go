package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strconv"
  "strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
  return false
}

func main() {
  file_name := "file_test.txt"

  file, err := os.Open(file_name)
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  var tab []string

  for scanner.Scan() {
    tab = append(tab, scanner.Text())
  }

  fmt.Println(tab)
  fmt.Println(tab[0])
  fmt.Println(len(tab))
  for i := 0; i < len(tab); i++{
    fmt.Println(strings.Split(tab[i], " "))
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  all_nods := make(map[string]map[string]int)
  var all_points []string

  for i := 0; i < len(tab); i++{
    tab_nod := strings.Split(tab[i], " ")
    b := make(map[string]int)
    if contains(all_points, tab_nod[0]) == false{
      all_points = append(all_points, tab_nod[0])
    }
    if contains(all_points, tab_nod[1]) == false{
      all_points = append(all_points, tab_nod[1])
    }
    dist, _ := strconv.Atoi(tab_nod[2])
    if val, check := all_nods[tab_nod[0]]; check {
      b = val
      b[tab_nod[1]] = dist
      all_nods[tab_nod[0]] = b
    } else {
      b[tab_nod[1]] = dist
      all_nods[tab_nod[0]] = b
    }
  }

  fmt.Println(all_points)
  fmt.Println(all_nods["B"]["C"])
  fmt.Println(len(all_nods["B"]))

  fmt.Println("De quel point commecer ? ")
  var start_point string
  fmt.Scanln(&start_point)
  fmt.Println("Point d'arrêt ? ")
  var end_point string
  fmt.Scanln(&end_point)
  fmt.Println(start_point, end_point)
  res := start_point
  next_point := start_point

  for next_point != end_point {
    if len(all_nods[next_point]) != 1 {
      dist := 100000
      var new_point string
      for i := 0; i < len(all_points); i++ {
        if _, check := all_nods[next_point][all_points[i]]; check {
          if dist > all_nods[next_point][all_points[i]] {
            dist = all_nods[next_point][all_points[i]]
            new_point = all_points[i]
          }
        }
      }
      next_point = new_point
    } else {
      for i := 0; i < len(all_points); i++ {
        if _, check := all_nods[next_point][all_points[i]]; check {
          next_point = all_points[i]
          break
        }
      }
    }
    res = res + " --> " + next_point
  }
  fmt.Println("Le chemin le plus court de", start_point, "à", end_point, "est le suivant :", res)
}
