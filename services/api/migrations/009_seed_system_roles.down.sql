DELETE FROM roles
WHERE organization_id IS NULL
  AND name IN (
      'Organization Admin',
      'HR Manager',
      'Recruiter',
      'Interviewer'
  );