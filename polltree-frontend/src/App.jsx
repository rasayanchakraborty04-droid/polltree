import { BrowserRouter, Routes, Route } from "react-router-dom"

import Home from "./pages/Home"
import Login from "./pages/Login"
import PageNotExist from "./components/PageNotExist"

import Navbar from "./components/Navbar"

export default function App() {
  return (
    <>
      <Navbar />

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/ppe" element={<PageNotExist />} />
      </Routes>
    </>
  )
}