CREATE TABLE go-ebuild_cache(
   "URL"           VARCHAR(30) NOT NULL PRIMARY KEY
  ,"gopkgname"     VARCHAR(30) NOT NULL
  ,"gentoopkgname" VARCHAR(30) NOT NULL
  ,"EGO_SUM"       VARCHAR(3000) NOT NULL
);
INSERT INTO go-ebuild_cache("URL","gopkgname","gentoopkgname","EGO_SUM") VALUES (NULL,NULL,NULL,NULL);
