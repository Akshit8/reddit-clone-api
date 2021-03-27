CREATE TABLE post (
    id SERIAL NOT NULL, 
    title character varying NOT NULL,
    createdAt TIMESTAMP NOT NULL DEFAULT now(), 
    updatedAt TIMESTAMP NOT NULL DEFAULT now()
);