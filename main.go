package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func main() {
	app := fiber.New()

	products := []Product{
		{1, "Cabe", 62000, 120},
		{2, "Daging ayam", 25000, 150},
		{3, "Gula", 20000, 200},
	}

	// menampilkan semua produk
	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(products)
	})

	//menampilkan 1 produk
	app.Get("/products/:id", func(c *fiber.Ctx) error {
		paramID := c.Params("id")
		id, _ := strconv.Atoi(paramID)

		var foundProduct Product
		for _, p := range products {
			if p.ID == id {
				foundProduct = p
				break
			}
		}
		return c.JSON(foundProduct)
	})

	//menambahkan produk
	app.Post("/product", func(c *fiber.Ctx) error {
		var newProduct Product
		if err := c.BodyParser(&newProduct); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
			return err
		}

		newProduct.ID = len(products) + 1

		products = append(products, newProduct)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Product created successfully",
			"product": newProduct,
		})

	})

	//memperbaharui data produk
	app.Put("/product/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var updateProduct Product
		if err := c.BodyParser(&updateProduct); err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid Request Body",
			})
			return err
		}

		var foundIndex int = -1
		for i, p := range products {
			if p.ID == id {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {
			products[foundIndex] = updateProduct
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d update successfully", id),
				"product": updateProduct,
			})
		} else {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d not found", id),
			})
		}
	})

	//delete product
	app.Delete("/product/:id", func(c *fiber.Ctx) error {

		id, _ := strconv.Atoi(c.Params("id"))

		var foundIndex int = -1
		for i, p := range products {
			if p.ID == id {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {

			products = append(products[:foundIndex], products[foundIndex+1:]...)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
			})
		} else {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d not found", id),
			})
		}
	})

	app.Listen(":3000")
}
