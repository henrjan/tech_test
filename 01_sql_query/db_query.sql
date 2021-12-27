SELECT usr1.ID, usr1.UserName, usr2.ParentUserName
FROM USER usr1
OUTER JOIN (
    Select ID, UserName as ParentUserName
    FROM USER
) usr2 ON usr1.Parent = usr2.ID
ORDER BY usr1.ID;