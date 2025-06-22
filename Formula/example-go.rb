class ExampleGo < Formula
  desc "A brief description of your application"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/example-go-d1fd30bb1a8b5048588dc8b637c63e6c3adc8c2a/example-go"
  version "d1fd30bb1a8b5048588dc8b637c63e6c3adc8c2a"

  def install
    bin.install "example-go"

#     system bin/"example-go", "man"
#     man1.install Dir["*.1"]
#
#     generate_completions_from_executable(bin/"example-go", "completion")
  end
end
