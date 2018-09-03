magic_number <- c(505.0,279.3,540.5,202.4,160.0,388.8,446.3,560.0,492.9,593.0,172.8,316.8,504.1,490.2,672.8,592.9,540.0,494.7,689.7,168.3)
sample_iris <- iris[1:20,]

# fit linear model
model <- lm(sample_iris$Sepal.Length ~ magic_number - 1)
pred <- predict(model, data.frame(sample_iris$Sepal.Length))

# magic!
print(rawToChar(as.raw(pred / model$coefficients / sample_iris$Sepal.Length)))
