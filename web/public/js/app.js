let auth0 = null;

const fetchAuthConfig = () => fetch("/auth_config.json");

const configureClient = async() => {
    const response = await fetchAuthConfig();
    const config = await response.json();

    auth0 = await createAuth0Client({
        domain: config.domain,
        client_id: config.clientId
    });
};

window.onload = async() => {
    await configureClient();
    updateUI();

    const isAuthenticated = await auth0.isAuthenticated();

    if (isAuthenticated) {
        // show the gated content
        return;
    }

    // Check for the code and state parameters
    const query = window.location.search;
    if (query.includes("code=") && query.includes("state=")) {

        // Process the login state
        await auth0.handleRedirectCallback();

        updateUI();

        // Use replaceState to redirect the user away and remove the querystring parameters
        window.history.replaceState({}, document.title, "/");
    }

    homeAnchor = document.getElementById("home-anchor");
    homeAnchor.addEventListener('click', renderHome);
}

const updateUI = async() => {
    const isAuthenticated = await auth0.isAuthenticated();

    if (isAuthenticated) {
        document.getElementById("session-btn").innerText = "Log out";
    } else {
        document.getElementById("profile").classList.add("hidden");
        document.getElementById("session-btn").innerText = "Log in";
    }
};

const changeStateButton = async() => {
    const isAuthenticated = await auth0.isAuthenticated();
    if (isAuthenticated) {
        document.getElementById("session-btn").innerText = "Log out";
        await logout()
    } else {
        document.getElementById("session-btn").innerText = "Log in";
        await login()
    }
}

const login = async() => {
    await auth0.loginWithRedirect({
        redirect_uri: window.location.origin
    });
};

const logout = () => {
    auth0.logout({
        returnTo: window.location.origin
    });
};

const renderHome = () => {
    hideProfile();

};

const renderProfile = async() => {
    var profile = await auth0.getUser();
    var user = {
        "username": profile.nickname,
        "email": profile.email,
        "picture": profile.picture,
        "subject": profile.sub,
    };
    axios({
            method: 'post',
            url: 'http://localhost:8080/api/v1/users/',
            data: user
        })
        .then((response) => {
            console.log(response);
        }, (error) => {
            console.log(error);
        });

    document.getElementById("profile").classList.remove("hidden");
    document.getElementById("username").textContent = user.username;
    document.getElementById("profile-picture").setAttribute("src", user.picture);
};

const hideProfile = async() => {
    document.getElementById("profile").classList.add("hidden");
}