import { useEffect, useState } from "react";
import api from "./services/api";

const App = () => {
  const [status, setStatus] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await api.get("/health");
        setStatus(response.data.status);
      } catch (error) {
        console.error(error);
        setStatus("Error");
      }
    }
    fetchData()
  }, [])

  return <div>Backend Status: {status || "Loading..."}</div>
}

export default App;
