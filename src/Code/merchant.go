package gokemon

import (
	"fmt"
	"strings"
)

func VisiteMarchand(joueur *Dresseur) {
	for {
		largeur := 155
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		AfficherTitre()

		fmt.Println(Jaune("╔" + strings.Repeat("═", largeur-2) + "╗"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("                                                           BOUTIQUE DU MARCHAND", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu(fmt.Sprintf("  Votre Porte-Monnaie: %d PokéDollars", joueur.Argent), largeur)
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu("  Capacité d'inventaire: "+fmt.Sprintf("%d/%d", len(joueur.Inventaire), joueur.CapaciteInventaire), largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("", largeur)
		AfficherLigneMenu(" Votre inventaire :", largeur)
		for i, item := range joueur.Inventaire {
			AfficherLigneMenu(fmt.Sprintf("%d. %s (x%d)", i+1, item.Nom, item.Quantite), largeur)
		}
		fmt.Println(Jaune("╠" + strings.Repeat("═", largeur-2) + "╣"))
		AfficherLigneMenu("1. Acheter une Potion de Soin (50 PokéDollars)", largeur)
		AfficherLigneMenu("2. Acheter une Pokéball (100 PokéDollars)", largeur)
		AfficherLigneMenu("3. Acheter une Potion de Poison (75 PokéDollars)", largeur)
		AfficherLigneMenu("4. Acheter une Augmentation d'inventaire (30 PokéDollars)", largeur)
		AfficherLigneMenu("5. Vendre un objet", largeur)
		AfficherLigneMenu("6. Vendre un Pokémon", largeur)
		AfficherLigneMenu("7. Retour au menu principal", largeur)
		AfficherLigneMenu("", largeur)
		fmt.Println(Jaune("╚" + strings.Repeat("═", largeur-2) + "╝"))

		fmt.Printf(Jaune("\nVotre Porte-Monnaie: %d PokéDollars\n"), joueur.Argent)
		fmt.Print(Vert("\nEntrez votre choix (1-7): "))
		var choix string
		Wrap(func() { fmt.Scanln(&choix) })

		switch choix {
		case "1":
			AcheterObjet(joueur, "Potion de Soin", 50)
		case "2":
			AcheterObjet(joueur, "Pokéball", 100)
		case "3":
			AcheterObjet(joueur, "Potion de Poison", 75)
		case "4":
			AcheterAugmentationInventaire(joueur)
		case "5":
			VendreObjet(joueur)
		case "6":
			VendrePokemon(joueur)
		case "7":
			return
		default:
			fmt.Println(Jaune("\nChoix invalide. Veuillez réessayer."))
		}

		fmt.Print(Vert("\nAppuyez sur Entrée pour continuer..."))
		Wrap(func() { fmt.Scanln() })
	}
}

func AcheterObjet(joueur *Dresseur, nomObjet string, prix int) {
	if joueur.Argent >= prix {
		totalItems := 0
		for _, item := range joueur.Inventaire {
			totalItems += item.Quantite
		}
		if totalItems >= joueur.CapaciteInventaire {
			fmt.Println(Jaune("\nVotre inventaire est plein. Vous ne pouvez pas acheter plus d'objets."))
			return
		}

		joueur.Argent -= prix
		for i := range joueur.Inventaire {
			if joueur.Inventaire[i].Nom == nomObjet {
				joueur.Inventaire[i].Quantite++
				fmt.Printf(Jaune("\nVous avez acheté un(e) %s pour %d PokéDollars.\n"), nomObjet, prix)
				return
			}
		}
		ajouterObjetInventaire(joueur, nomObjet, 1)
		fmt.Printf(Jaune("\nVous avez acheté un(e) %s pour %d PokéDollars.\n"), nomObjet, prix)
	} else {
		fmt.Println(Jaune("\nVous n'avez pas assez d'argent pour acheter cet objet."))
	}
}

func VendreObjet(joueur *Dresseur) {
	fmt.Println(Jaune("\nQuels objets voulez-vous vendre ?"))
	for i, item := range joueur.Inventaire {
		fmt.Printf(Jaune("%d. %s (x%d) - Prix de vente: %d PokéDollars\n"), i+1, item.Nom, item.Quantite, GetPrixVente(item.Nom))
	}
	fmt.Printf(Jaune("%d. Annuler\n"), len(joueur.Inventaire)+1)

	var choix int
	fmt.Print(Vert("\nEntrez votre choix : "))
	Wrap(func() { fmt.Scanln(&choix) })

	if choix > 0 && choix <= len(joueur.Inventaire) {
		item := &joueur.Inventaire[choix-1]
		if item.Quantite > 0 {
			prixVente := GetPrixVente(item.Nom)
			joueur.Argent += prixVente
			item.Quantite--
			fmt.Printf(Jaune("\nVous avez vendu un(e) %s pour %d PokéDollars.\n"), item.Nom, prixVente)
			if item.Quantite == 0 {
				joueur.Inventaire = append(joueur.Inventaire[:choix-1], joueur.Inventaire[choix:]...)
			}
		} else {
			fmt.Println(Jaune("\nVous n'avez plus de cet objet dans votre inventaire."))
		}
	} else if choix != len(joueur.Inventaire)+1 {
		fmt.Println(Jaune("\nChoix invalide."))
	}
}

func VendrePokemon(joueur *Dresseur) {
	if len(joueur.Equipe) <= 1 {
		fmt.Println(Jaune("\nVous ne pouvez pas vendre votre dernier Pokémon !"))
		return
	}

	fmt.Println(Jaune("\nQuels Pokémon voulez-vous vendre ?"))
	for i, pokemon := range joueur.Equipe {
		prix := pokemon.Niveau * 100
		fmt.Printf(Jaune("%d. %s (Niveau %d) - Prix de vente: %d PokéDollars\n"), i+1, pokemon.Nom, pokemon.Niveau, prix)
	}
	fmt.Printf(Jaune("%d. Annuler\n"), len(joueur.Equipe)+1)

	var choix int
	fmt.Print(Vert("\nEntrez votre choix : "))
	Wrap(func() { fmt.Scanln(&choix) })

	if choix > 0 && choix < len(joueur.Equipe)+1 {
		pokemonVendu := joueur.Equipe[choix-1]
		prixVente := pokemonVendu.Niveau * 100
		joueur.Argent += prixVente
		joueur.Equipe = append(joueur.Equipe[:choix-1], joueur.Equipe[choix:]...)
		fmt.Printf(Jaune("\nVous avez vendu %s pour %d PokéDollars.\n"), pokemonVendu.Nom, prixVente)
	} else if choix != len(joueur.Equipe)+1 {
		fmt.Println(Jaune("\nChoix invalide."))
	}
}

func GetPrixVente(nomObjet string) int {
	switch nomObjet {
	case "Potion":
		return 25
	case "Pokéball":
		return 50
	case "Potion de Poison":
		return 150
	default:
		return 35
	}
}

func AcheterAugmentationInventaire(joueur *Dresseur) {
	if joueur.Argent < 30 {
		fmt.Println(Jaune("\nVous n'avez pas assez d'argent pour acheter une Augmentation d'inventaire."))
		return
	}

	if joueur.CapaciteInventaire >= MaxAugmentationsInv {
		fmt.Println(Jaune("\nVous avez déjà atteint le nombre maximum d'augmentations d'inventaire."))
		return
	}

	joueur.Argent -= 30
	upgradeInventorySlot(joueur)
	fmt.Println(Jaune("\nVous avez acheté une Augmentation d'inventaire pour 30 PokéDollars."))
}

func ajouterObjetInventaire(joueur *Dresseur, nomObjet string, quantite int) {
	for i := range joueur.Inventaire {
		if joueur.Inventaire[i].Nom == nomObjet {
			joueur.Inventaire[i].Quantite += quantite
			return
		}
	}
	joueur.Inventaire = append(joueur.Inventaire, InventoryItem{Nom: nomObjet, Quantite: quantite})
}
