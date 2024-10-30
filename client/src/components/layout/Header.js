import React from "react";
import { Container } from "react-bootstrap";
import { Link } from "react-router-dom";
import { useSelector } from "react-redux";

const Header = () => {
  const { loggedIn, user } = useSelector((state) => state.auth);

  return (
    <>
      <Container fluid className="container-fluid header">
        <h1 className="text-center text-uppercase">
          Welcome to Web Blog
        </h1>
      </Container>
      <Container>
        <nav>
          <ul className="menu">
            <li>
              <Link to="/" aria-label="Home">Home</Link>
            </li>
            <li>
              <Link to="/blog" aria-label="Blog">Blog</Link>
            </li>
            <li>
              <Link to="/about" aria-label="About">About</Link>
            </li>
            <li>
              <Link to="/contact" aria-label="Contact">Contact</Link>
            </li>
            <li>
              {loggedIn ? (
                <div className="user-info">
                  Welcome back, {user.user_id}
                  <Link to="/logout" aria-label="Logout"> &nbsp; &nbsp; Logout</Link>
                </div>
              ) : (
                <Link to="/login" aria-label="Login">Login</Link>
              )}
            </li>
          </ul>
        </nav>
      </Container>
    </>
  );
};

export default Header;
