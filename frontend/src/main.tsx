import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './index.css'
import App from './App.tsx'
import Login from './pages/login/Login.tsx'
import Register from './pages/register/Register.tsx'
import Header from './components/custom-ui/header.tsx'
import Footer from './components/custom-ui/footer.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Header />
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<App />} />
        <Route path='login' element={<Login />} />
        <Route path='register' element={<Register />} />
      </Routes>
    </BrowserRouter>
    <Footer />
  </StrictMode>,
)
