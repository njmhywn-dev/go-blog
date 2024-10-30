/* eslint-disable react/jsx-no-target-blank */
import React from "react";
import { Col, Container, Row } from "react-bootstrap";
import { Link } from "react-router-dom";

const Footer = () => {
  return (
    <Container fluid className="container-fluid footer">
      <Row>
        <Col xs="6">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/blog">Blog</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/contact">Contact</Link>
            </li>
          </ul>
        </Col>
        <Col xs="6">
          <div className="col text-center text-lg-left pt-3">
            <a className="" href="/contact/">
              <h6>Connect With Us</h6>
            </a>
            <div className="social-links pt-3">
              <a
                target="_blank"
                href="https://www.youtube.com"
                className="youtube"
                title="Connect on Youtube"
              >
                <i className="fab fa-youtube fa-3x"></i>
              </a>
              <a
                target="_blank"
                href="https://github.com/njmhywn-dev"
                className="Github"
                title="Connect on Github"
              >
                <i className="fab fa-github fa-3x"></i>
              </a>
            </div>
          </div>
        </Col>
      </Row>
    </Container>
  );
};

export default Footer;