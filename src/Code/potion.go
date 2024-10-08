package gokemon

import (
	"fmt"
	"time"
)

func UsePotion(joueur *Dresseur, pokemon *Pokemon) {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Potion de Soin" {
			if item.Quantite > 0 {
				if pokemon.PVActuels == pokemon.PVMax {
					fmt.Printf(Jaune("\n%s a déjà tous ses PV.\n"), pokemon.Nom)
					return
				}
				healAmount := 15
				pokemon.PVActuels += healAmount
				if pokemon.PVActuels > pokemon.PVMax {
					pokemon.PVActuels = pokemon.PVMax
				}
				joueur.Inventaire[i].Quantite--
				fmt.Printf(Jaune("\nVous avez utilisé une Potion. %s a récupéré %d PV. "), pokemon.Nom, healAmount)
				fmt.Println(Jaune(afficherBarrePV(*pokemon)))
			} else {
				fmt.Println(Jaune("\nVous n'avez plus de Potions."))
			}
			return
		}
	}
	fmt.Println(Jaune("\nVous n'avez pas de Potions dans votre inventaire."))
}

func PoisonPot(enemy *Pokemon) {
	fmt.Printf(Jaune("\nLa Potion de Poison affecte %s !\n"), enemy.Nom)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		enemy.PVActuels -= 7
		if enemy.PVActuels < 0 {
			enemy.PVActuels = 0
		}
		fmt.Printf(Jaune("%s subit 7 points de dégâts."), enemy.Nom)
		fmt.Println(Jaune(afficherBarrePV(*enemy)))
		if enemy.PVActuels == 0 {
			fmt.Printf(Jaune("%s est K.O. !\n"), enemy.Nom)
			break
		}
	}
}

func UsePoisonPotion(joueur *Dresseur, pokemon *Pokemon) bool {
	for i, item := range joueur.Inventaire {
		if item.Nom == "Potion de Poison" {
			if item.Quantite > 0 {
				joueur.Inventaire[i].Quantite--
				PoisonPot(pokemon)
				return true
			} else {
				fmt.Println(Jaune("\nVous n'avez plus de Potions de Poison."))
				return false
			}
		}
	}
	fmt.Println(Jaune("\nVous n'avez pas de Potions de Poison dans votre inventaire."))
	return false
}
