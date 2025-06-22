class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-391d7b808e026a7c010ffd4fd2f720c8991c5eb1/example-go"
  version "391d7b808e026a7c010ffd4fd2f720c8991c5eb1"

  def install
    bin.install "example-go"

    system "./example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable("./example-go", "completion")
  end
end
