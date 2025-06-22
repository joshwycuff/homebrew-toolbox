class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-df58f9ebcc96ccb03099632a4666c1cdb555a7ac/example-go"
  version "df58f9ebcc96ccb03099632a4666c1cdb555a7ac"

  def install
    bin.install "example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
