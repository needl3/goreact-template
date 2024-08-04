import { createBrowserRouter } from "react-router-dom"
import { HomeController } from "./controllers/home"
import { Layout } from "./pages/layout"
import { NewController } from "./controllers/new"

export const router = createBrowserRouter([
  { path: "/", element: <HomeController /> },
  {
    path: "/app", element: <Layout />, children: [
      { path: "new", element: <NewController /> }
    ]
  }
])

