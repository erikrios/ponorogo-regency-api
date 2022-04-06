-- Provinces --

-- Get all provinces
SELECT p.id, p.name
FROM provinces p;

-- Get a province by id
SELECT p.id, p.name
FROM provinces p
WHERE p.id = '35';

-- Search provinces by name
SELECT p.id, p.name
FROM provinces p
WHERE p.name ILIKE '%Jawa%';

-- Regencies --

-- Get all regencies
SELECT r.id, r.name, r.province_id, p.name AS province_name
FROM regencies r
         INNER JOIN provinces p on r.province_id = p.id;

-- Get a regencies by id
SELECT r.id, r.name, r.province_id, p.name AS province_name
FROM regencies r
         INNER JOIN provinces p on r.province_id = p.id
WHERE r.id = '3502';

-- Search regencies by name
SELECT r.id, r.name, r.province_id, p.name AS province_name
FROM regencies r
         INNER JOIN provinces p on r.province_id = p.id
WHERE r.name ILIKE '%ROG%';

-- Districts --

-- Get all districts
SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name
FROM districts d
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id;

-- Get a district by id
SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name
FROM districts d
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE d.id = '3502030';

-- Search districts by name
SELECT d.id, d.name, d.regency_id, r.name AS regency_name, r.province_id, p.name AS province_name
FROM districts d
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE d.name ILIKE '%uNg%';

-- Villages --

-- Get all villages
SELECT v.id,
       v.name,
       v.district_id,
       d.name AS district_name,
       d.regency_id,
       r.name AS regency_name,
       r.province_id,
       p.name AS province_name
FROM villages v
         INNER JOIN districts d on d.id = v.district_id
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id;

-- Get a village by id
SELECT v.id,
       v.name,
       v.district_id,
       d.name AS district_name,
       d.regency_id,
       r.name AS regency_name,
       r.province_id,
       p.name AS province_name
FROM villages v
         INNER JOIN districts d on d.id = v.district_id
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE v.id = '3502030007';

-- Search villages by name
SELECT v.id,
       v.name,
       v.district_id,
       d.name AS district_name,
       d.regency_id,
       r.name AS regency_name,
       r.province_id,
       p.name AS province_name
FROM villages v
         INNER JOIN districts d on d.id = v.district_id
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE v.name ILIKE '%aGE%';

-- Get villages by district id
SELECT v.id,
       v.name,
       v.district_id,
       d.name AS district_name,
       d.regency_id,
       r.name AS regency_name,
       r.province_id,
       p.name AS province_name
FROM villages v
         INNER JOIN districts d on d.id = v.district_id
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE v.district_id = '3502030';

-- Get villages by district name
SELECT v.id,
       v.name,
       v.district_id,
       d.name AS district_name,
       d.regency_id,
       r.name AS regency_name,
       r.province_id,
       p.name AS province_name
FROM villages v
         INNER JOIN districts d on d.id = v.district_id
         INNER JOIN regencies r on d.regency_id = r.id
         INNER JOIN provinces p on r.province_id = p.id
WHERE d.name ILIKE '%uNg%';
