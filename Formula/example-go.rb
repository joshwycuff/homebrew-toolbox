class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-40056adf7ed6546ef5d64bb1072e09b4d1cc0d56/example-go"
  version "40056adf7ed6546ef5d64bb1072e09b4d1cc0d56"

  def install
    bin.install "example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
