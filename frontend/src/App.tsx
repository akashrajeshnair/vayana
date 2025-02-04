import AppRoutes from "./routes/routes";
import Header from "./components/custom-ui/header";
import Footer from "./components/custom-ui/footer";

const App = () => {
  return (
    <>
      <Header />
      <AppRoutes />
      <Footer />
    </>
  )
}

export default App;
