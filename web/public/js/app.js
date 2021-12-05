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

    // NEW - check for the code and state parameters
    const query = window.location.search;
    if (query.includes("code=") && query.includes("state=")) {

        // Process the login state
        await auth0.handleRedirectCallback();

        updateUI();

        // Use replaceState to redirect the user away and remove the querystring parameters
        window.history.replaceState({}, document.title, "/");
    }
}

const updateUI = async() => {
    const isAuthenticated = await auth0.isAuthenticated();

    if (isAuthenticated) {
        document.getElementById("profile").classList.remove("hidden");

        document.getElementById(
            "ipt-access-token"
        ).innerHTML = await auth0.getTokenSilently();

        document.getElementById("ipt-user-profile").textContent = JSON.stringify(
            await auth0.getUser()
        );
        document.getElementById("btn-login").innerText = "Log out";
        renderProfile(await auth0.getUser())
    } else {
        document.getElementById("profile").classList.add("hidden");
        document.getElementById("btn-login").innerText = "Log in";
    }
};

const changeStateButton = async() => {
    const isAuthenticated = await auth0.isAuthenticated();
    if (isAuthenticated) {
        document.getElementById("btn-login").innerText = "Log out";
        await logout()
    } else {
        document.getElementById("btn-login").innerText = "Log in";
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

const renderProfile = (profile) => {
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
}