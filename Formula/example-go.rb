class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-b185f0ccd960959ee20ad2c2288a06d01da81191/example-go"
  version "b185f0ccd960959ee20ad2c2288a06d01da81191"
  
  depends_on "go"

  def install
    bin.install "example-go"

    system bin/"example-go", "man"
    man1.install Dir["*.1"]

    generate_completions_from_executable(bin/"example-go", "completion")
  end
end
