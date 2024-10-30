/* eslint-disable jsx-a11y/alt-text */
import { useEffect, useState } from "react";
import { Container, Row, Col, Card, Button } from "react-bootstrap"
import axios from "axios"
import "../App.css"
import { Link } from "react-router-dom";
import Spinner from "react-bootstrap/Spinner";

const Home = () => {
    const [apiData, setApiData] = useState(false)
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const fetchData = async () => {

            try {
                const apiUrl = process.env.REACT_APP_API_ROOT;
                const response = await axios.get(apiUrl)

                if (response.status === 200) {
                    if (response?.data.statusText === "Ok") {
                        setApiData(response?.data?.blog_records
                        )
                    }
                }

                setLoading(false);
            } catch (error) {
                setLoading(false);
                console.log(error.response)
            }
        }

        fetchData()
        return () => {

        }
    }, [])

    console.log(apiData)

    if (loading) {
        return (
            <>
                <Container className="spinner">
                    <Spinner animation="border" variant="primary" />
                </Container>
            </>
        );
    }

    return (
        <Container className="py-2">


            <Row className="justify-content-md-center">

                <h3>
                    <Link to="add" className="btn btn-primary">
                        add blog +
                    </Link>
                </h3>


                {apiData &&
                    apiData.map((record, index) => (
                        <Col key={index} xs="3" className="py-6">
                            <Card style={{ width: '18rem' }}>
                                <Card.Img variant="top" src={`${process.env.REACT_APP_API_ROOT_IMG}/${record.image}`} />
                                <Card.Body>
                                    <Card.Title><Link to={`/blog/${record.id}`}>{record.title}</Link></Card.Title>
                                    <Card.Text>
                                        Some quick example text to build on the card title and make up the
                                        bulk of the card's content.
                                    </Card.Text>
                                    <Button variant="primary">
                                        <Link to={`edit/${record.id}`}>
                                            Edit
                                        </Link>
                                    </Button>
                                    {' '}
                                    <Button variant="danger">
                                        <Link to={`delete/${record.id}`}>
                                            Delete
                                        </Link>
                                    </Button>
                                </Card.Body>
                            </Card>
                        </Col>
                    ))}
            </Row>
        </Container>
    )
}

export default Home;