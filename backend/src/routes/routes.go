package routes

import (
	"shared-blog-backend/src/controllers"
	"shared-blog-backend/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	
	// Public Routes
	publicApi := app.Group("/api")

	// api.Get("/user-posts", controllers.UserPosts)
	publicApi.Post("/user/signup", controllers.SignUp)
	publicApi.Post("/user/signin", controllers.SignIn)
	publicApi.Post("/user/signout", controllers.SignOut)
	publicApi.Post("/user/is-authorized", controllers.IsAuthorized)
	
	publicApi.Get("/post/get-posts", controllers.GetPosts)
	publicApi.Get("/post/get-post/:postId", controllers.GetPost)


	// Protected Routes
	protectedApi := app.Group("/api/protected", middleware.ProtectedRoute())
	
	protectedApi.Post("/post/create-post", controllers.CreatePost)
	protectedApi.Post("/post/vote/:postId", controllers.Vote)
	protectedApi.Post("/post/unvote/:postId", controllers.UnVote)

	protectedApi.Get("/topics/get-topics", controllers.GetTopics)

}
