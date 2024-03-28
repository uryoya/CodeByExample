SELECT
  `main`.`Media`.`id`,
  `main`.`Media`.`url`
FROM
  `main`.`Media`
WHERE (
  NOT (`main`.`Media`.`id`)
  NOT IN (
    SELECT
      `main`.`Image`.`id`
    FROM
      `main`.`Image`
    WHERE
      `main`.`Image`.`id`
    IS NOT NULL
  )
)
ORDER BY `main`.`Media`.`title` ASC
LIMIT ? OFFSET ?
