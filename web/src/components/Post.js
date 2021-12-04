import React from "react";
import { Row, Col } from "reactstrap";
import contentData from "../utils/contentData";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import logo from "../assets/logo.svg";

const Post = () => (
	<div className="next-steps my-5">
		{contentData.map((col, i) => (
			<Row className="d-flex justify-content-between">
				<h6 className="mb-3">
					<a href={col.link}>
						<img className="mb-3 app-logo" src={logo} alt="User profile picture" width="50" />
						<FontAwesomeIcon icon="link" className="mr-2" />
						{col.title}
					</a>
				</h6>
				<p>{col.description}</p>
			</Row>
		))}
	</div>
);

export default Post;
