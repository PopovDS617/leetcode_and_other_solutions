const validateEmailAddress = (email) => {
  const regex = /[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]+/gi;

  const isValidEmail = regex.test(email);

  return isValidEmail;
};
