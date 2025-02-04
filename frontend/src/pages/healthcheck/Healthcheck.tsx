import { useState, useEffect } from "react";
import api from "@/services/api";
import { useSelector } from "react-redux";
import { store } from "@/store/store";

const Healthcheck = () => {
  const [status, setStatus] = useState("");
  const isAuthenticated = useSelector((state: RootState) => state.auth.isAuthenticated)

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

  return (
    <>
      <div>Backend Status: {status || "Loading..."}</div>
      <div>Auth status: {isAuthenticated}</div>
    </>
  )
}

export default Healthcheck;
