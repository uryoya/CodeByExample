lazy val root = (project in file(".")).
  settings(
    name := "library-examples",
    scalaVersion := "2.12.5",
    libraryDependencies ++= Seq(
      "co.fs2" %% "fs2-core" % "0.10.4",
    )
  )
