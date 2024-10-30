import { Routes, Route } from "react-router-dom";
import Header from "./components/layout/Header";
import Home from "./page/Home";
import Blog from "./page/Blog";
import Add from "./page/Add";
import Edit from "./page/Edit";
//import Footer from "./components/layout/Footer";
import "./App.css";
import Delete from "./page/Delete";
import Login from "./page/Login";
import Logout from "./page/Logout";
import About from "./page/About";
import Contact from "./page/Contact";
import { useEffect } from "react";
import axios from "axios";


function App() {

  const token = window.localStorage.getItem("token");

  useEffect(() => {
    const timestamp = 1000 * 60 * 3 - 5;
    // const timestamp = 10000;

    let interval = setInterval(() => {
      if (token !== null) {
        updateToken();
      }
    }, timestamp);

    return () => {
      clearInterval(interval);
    };
  }, [token]);

  const updateToken = async () => {
    try {
      const apiUrl = `${process.env.REACT_APP_AUTH_API}/private/refreshtoken`;

      const response = await axios.get(apiUrl, {
        headers: {
          token: window.localStorage.getItem("token"),
        },
      });

      if (response.status === 200) {
        const data = await response.data;

        window.localStorage.setItem("token", data.token);
      }
    } catch (error) {
      console.log(error);

      window.localStorage.removeItem("token");
    }

    console.log("Inside update token");
  };

  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<><h1 className="text-center" >Welcome to My Web Site</h1></>} />
        <Route path="/blog" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/blog/Add" element={<Add />} />
        <Route path="/blog/Edit/:id" element={<Edit />} />
        <Route path="/blog/Delete/:id" element={<Delete />} />
        <Route path="/login" element={<Login />} />
        <Route path="/logout" element={<Logout />} />
        <Route path="/about" element={<About />} />
        <Route path="/contact" element={<Contact />} />
      </Routes>
    </>
  );
}

export default App;
