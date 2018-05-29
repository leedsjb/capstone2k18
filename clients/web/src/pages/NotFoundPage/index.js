import React from "react";
import { Helmet } from "react-helmet";

const NotFoundPage = () => {
    return (
        <div>
            <Helmet>
                <title>Page not found</title>
            </Helmet>
            <div>Page not found</div>
        </div>
    );
};

export default NotFoundPage;
