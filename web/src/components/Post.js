import React from "react";
import logo from "../assets/logo.svg";

const Post = () => (
	<div class="container">
		<div class="row">
			<div class="col-sm-1">
				<img className="mb-3 app-logo" src={logo} alt="User profile picture" width="50" />
			</div>
			<div class="col-sm-6">
				<h5 className="mb-4">Username</h5>
			</div>
			<div class="col-sm-5">
				<h6 className="mb-4">Time</h6>
			</div>
		</div>
		<div class="row">
			<div class="col-sm">
				<p className="lead">
					Tweet
				</p>
			</div>
		</div>
	</div>
);

export default Post;
