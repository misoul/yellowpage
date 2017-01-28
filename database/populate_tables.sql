LOAD DATA LOCAL INFILE '/Users/vto/src/vto/yellowpage/database/comments.txt'
    INTO TABLE comments FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"'
    LINES TERMINATED BY '\n' IGNORE 1 LINES;
LOAD DATA LOCAL INFILE '/Users/vto/src/vto/yellowpage/database/companies.txt'
    INTO TABLE companies FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"'
    LINES TERMINATED BY '\n' IGNORE 1 LINES;
