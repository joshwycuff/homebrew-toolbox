class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-8541924f0e195adaf1f5ae90f8f72572e7e92712/example-go-8541924f0e195adaf1f5ae90f8f72572e7e92712.tar.gz"
  version "8541924f0e195adaf1f5ae90f8f72572e7e92712"

  def install
    bin.install "example-go"
    man1.install Dir["*.1"]
    generate_completions_from_executable("./example-go", "completion")
  end
end
